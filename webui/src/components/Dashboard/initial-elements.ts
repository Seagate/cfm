import { computed } from "vue";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useServiceStore } from "../Stores/ServiceStore";



export const useFlowData = () => {
  const applianceStore = useApplianceStore();
  const hostStore = useHostStore();
  const serviceStore = useServiceStore();

  const position = { x: 0, y: 0 };
  const serviceNodeType = "cfm-service";
  const applianceNodeType = "appliacne";
  const hostNodeType = "host";
  const bladeNodeType = "blade";

  const nodes = computed(() => {
    const coreNode = serviceStore.serviceVersion
      ? [
        {
          id: "cfm-service",
          label: "CFM Service",
          position: position,
          style: { backgroundColor: "#ffcc00", color: "#000" },
          type: serviceNodeType,
        },
      ]
      : [];

    const applianceNodes = applianceStore.applianceIds.flatMap(
      (appliance, applianceIndex) => [
        {
          id: `appliance-${appliance.id}`,
          label: appliance.id,
          position: position,
          style: { backgroundColor: "#00ccff", color: "#fff" },
          type: applianceNodeType,
        },
        ...appliance.bladeIds.map((bladeId, bladeIndex) => ({
          id: `blade-${bladeId}`,
          label: bladeId,
          position: position,
          style: { backgroundColor: "#ff6600", color: "#fff" },
          type: bladeNodeType,
        })),
      ]
    );

    const hostNodes = hostStore.hostIds.map((host, index) => ({
      id: `host-${host}`,
      label: host,
      position: position,
      style: { backgroundColor: "#66ff66", color: "#000" },
      type: hostNodeType,
    }));

    return [...coreNode, ...applianceNodes, ...hostNodes];
  });

  const edges = computed(() => {
    const coreEdges = serviceStore.serviceVersion
      ? applianceStore.applianceIds.flatMap((appliance) => [
        {
          id: `cfm-appliance-${appliance.id}`,
          type: "smoothstep",
          source: "cfm-service",
          target: `appliance-${appliance.id}`,
        },
        ...appliance.bladeIds.map((bladeId) => ({
          id: `appliance-blade-${appliance.id}-${bladeId}`,
          source: `appliance-${appliance.id}`,
          target: `blade-${bladeId}`,
        })),
      ])
      : [];

    const hostEdges = hostStore.hostIds.map((host) => ({
      id: `cfm-${host}`,
      source: "cfm-service",
      target: `host-${host}`,
    }));

    return [...coreEdges, ...hostEdges];
  });

  return { nodes, edges };
};
