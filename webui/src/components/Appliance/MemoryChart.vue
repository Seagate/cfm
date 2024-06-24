<template>
  <v-container
    class="chart-container"
    style="position: relative; display: flex; align-items: center; padding: 0"
  >
    <canvas ref="memoryCanvas"></canvas>
  </v-container>
</template>

<script lang="ts">
import { Chart, registerables } from "chart.js";
import { useBladeStore } from "../Stores/BladeStore";
import {
  computed,
  onMounted,
  ref,
  nextTick,
  defineComponent,
  watch,
} from "vue";
Chart.register(...registerables);

export default defineComponent({
  setup() {
    const bladeStore = useBladeStore();
    const selectedBladeTotalMemoryAvailableMiB = computed(
      () => bladeStore.selectedBladeTotalMemoryAvailableMiB
    );
    const selectedBladeTotalMemoryAllocatedMiB = computed(
      () => bladeStore.selectedBladeTotalMemoryAllocatedMiB
    );

    let totalMemoryMiB = 0;
    // Must handle the scenario where selectedBladeTotalMemoryAvailableMiB or selectedBladeTotalMemoryAllocatedMiB does not exist(value is 0).
    if (
      selectedBladeTotalMemoryAvailableMiB.value &&
      selectedBladeTotalMemoryAllocatedMiB.value
    ) {
      totalMemoryMiB =
        selectedBladeTotalMemoryAvailableMiB.value +
        selectedBladeTotalMemoryAllocatedMiB.value;
    } else if (selectedBladeTotalMemoryAvailableMiB.value) {
      totalMemoryMiB = selectedBladeTotalMemoryAvailableMiB.value;
    } else if (selectedBladeTotalMemoryAllocatedMiB.value) {
      totalMemoryMiB = selectedBladeTotalMemoryAllocatedMiB.value;
    }
    const totalMemoryGiB = totalMemoryMiB / 1024;

    const memoryCanvas = ref<HTMLCanvasElement | null>(null);

    let chartInstance: any = null; // Store a reference to the chart instance

    onMounted(async () => {
      await nextTick();
      if (memoryCanvas.value) {
        const ctx = memoryCanvas.value.getContext("2d");
        if (ctx) {
          if (
            selectedBladeTotalMemoryAvailableMiB.value !== undefined &&
            selectedBladeTotalMemoryAllocatedMiB.value !== undefined
          ) {
            chartInstance = new Chart(ctx, {
              type: "pie",
              data: {
                labels: ["Available Memory (GiB)", "Allocated Memory (GiB)"],
                datasets: [
                  {
                    data: [
                      selectedBladeTotalMemoryAvailableMiB.value / 1024,
                      selectedBladeTotalMemoryAllocatedMiB.value / 1024,
                    ],
                    backgroundColor: [
                      "rgba(110, 190, 74, 0.2)",
                      "rgba(255, 159, 64, 0.2)",
                    ],
                    borderColor: [
                      "rgba(110, 190, 74, 1)",
                      "rgba(255, 159, 64, 1)",
                    ],
                    borderWidth: 1,
                  },
                ],
              },
              options: {
                responsive: true,
                maintainAspectRatio: false,
                plugins: {
                  legend: {
                    position: "bottom",
                  },
                  title: {
                    display: true,
                    text: `Total Memory: ${totalMemoryGiB} GiB`,
                  },
                },
              },
            });
          }
        }
      }
    });

    // Watch for changes in blade memory
    watch(
      [
        selectedBladeTotalMemoryAvailableMiB,
        selectedBladeTotalMemoryAllocatedMiB,
      ],
      () => {
        if (chartInstance) {
          if (
            selectedBladeTotalMemoryAvailableMiB.value !== undefined &&
            selectedBladeTotalMemoryAllocatedMiB.value !== undefined
          ) {
            // Update the chart data
            chartInstance.data.datasets[0].data = [
              selectedBladeTotalMemoryAvailableMiB.value / 1024,
              selectedBladeTotalMemoryAllocatedMiB.value / 1024,
            ];

            // Update the chart
            chartInstance.update();
          }
        }
      }
    );

    return {
      memoryCanvas,
      totalMemoryGiB,
    };
  },
});
</script>
