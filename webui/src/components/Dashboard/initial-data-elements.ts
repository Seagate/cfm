// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { computed } from "vue";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useBladePortStore } from "../Stores/BladePortStore";

export const useData = () => {
    const applianceStore = useApplianceStore();
    const hostStore = useHostStore();
    const bladePortStore = useBladePortStore();

    const applianceNodeType = "appliance";
    const hostNodeType = "host";
    const bladeNodeType = "blade";

    const dataNodes = computed(() => {
        let currentYPosition = 0;
        const applianceNodes = applianceStore.applianceIds.flatMap(
            (appliance, index) => {
                const bladeCount = appliance.bladeIds.length;
                const applianceHeight = 50 + bladeCount * 50; // Adjust height based on number of blades
                const applianceWidth = 270; // Width of the appliance node
                const bladeWidth = 250; // Width of the blade node
                const bladeXPosition = (applianceWidth - bladeWidth) / 2; // Center the blade nodes

                const applianceNode = {
                    id: `appliance-${appliance.id}`,
                    data: { label: appliance.id, url: `/appliances/${appliance.id}` },
                    position: { x: 100, y: currentYPosition },
                    style: { backgroundColor: "rgba(242, 174, 114, 0.5)", color: "#000", height: `${applianceHeight}px`, width: `${applianceWidth}px` },
                    type: applianceNodeType,
                    sourcePosition: 'right',
                    targetPosition: 'left',
                };

                const bladeNodes = appliance.bladeIds.map((bladeId, bladeIndex) => ({
                    id: `blade-${bladeId}`,
                    data: { label: bladeId, url: `/appliances/${appliance.id}/blades/${bladeId}`, associatedAppliance: appliance.id },
                    position: { x: bladeXPosition, y: 50 + bladeIndex * 50 }, // Center blades within the appliance node
                    style: { backgroundColor: "#f2e394", color: "#000", width: `${bladeWidth}px` },
                    type: bladeNodeType,
                    parentNode: `appliance-${appliance.id}`,
                    extent: 'parent',
                    expandParent: true,
                    sourcePosition: 'right',
                    targetPosition: 'left',
                }));

                currentYPosition += applianceHeight + 20; // Add some space between nodes

                return [applianceNode, ...bladeNodes];
            }
        );

        const hostNodes = hostStore.hostIds.map((host, index) => ({
            id: `host-${host}`,
            data: { label: host, url: `/hosts/${host}` },
            position: { x: 500, y: index * 200 },
            style: { backgroundColor: "#d9ecd0", color: "#000" },
            type: hostNodeType,
            sourcePosition: 'right',
            targetPosition: 'left',
        }));

        return [...applianceNodes, ...hostNodes];
    });

    const dataEdges = computed(() => {
        const edges = bladePortStore.bladeIds.flatMap((blade) => [
            ...blade.connectedHostIds.map((hostId) => ({
                id: `appliance-blade-${blade.id}-${hostId}`,
                source: `blade-${blade.id}`,
                target: `host-${hostId}`,
                animated: true,
            }))
        ]);
        return edges;
    });

    return { dataNodes, dataEdges };
};
