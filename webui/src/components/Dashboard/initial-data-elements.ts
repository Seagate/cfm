// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { computed } from "vue";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useBladePortStore } from "../Stores/BladePortStore";
import { useLayout } from "./useLayout";

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
                const bladeCount = appliance.blades.length;
                const applianceHeight = 50 + bladeCount * 50; // Adjust height based on number of blades
                const applianceWidth = 270; // Width of the appliance node
                const bladeWidth = 250; // Width of the blade node
                const bladeXPosition = (applianceWidth - bladeWidth) / 2; // Center the blade nodes

                const applianceNode = {
                    id: `appliance-${appliance.id}`,
                    data: { label: appliance.id, url: `/appliances/${appliance.id}` },
                    position: { x: 100, y: currentYPosition },
                    style: { backgroundColor: useLayout().Colors.applianceColor, height: `${applianceHeight}px`, width: `${applianceWidth}px`, border: "none" },
                    type: applianceNodeType,
                    sourcePosition: 'right',
                    targetPosition: 'left',
                };

                const bladeNodes = appliance.blades.map((blade, bladeIndex) => {
                    const borderColor = useLayout().borderColorChange(blade.status);
                    return {
                        id: `blade-${blade.id}`,
                        data: { label: blade.id, url: `/appliances/${appliance.id}/blades/${blade.id}`, associatedAppliance: appliance.id },
                        position: { x: bladeXPosition, y: 50 + bladeIndex * 50 }, // Center blades within the appliance node
                        style: { backgroundColor: useLayout().Colors.baldeColor, width: `${bladeWidth}px`, border: `3px solid ${borderColor}` },
                        type: bladeNodeType,
                        parentNode: `appliance-${appliance.id}`,
                        extent: 'parent',
                        expandParent: true,
                        sourcePosition: 'right',
                        targetPosition: 'left',
                    }
                });

                currentYPosition += applianceHeight + 20; // Add some space between nodes

                return [applianceNode, ...bladeNodes];
            }
        );

        const hostNodes = hostStore.hostIds.map((host, index) => {
            const { width, height } = useLayout().measureText(host.id);
            const borderColor = useLayout().borderColorChange(host.status);

            return {
                id: `host-${host.id}`,
                data: { label: host.id, url: `/hosts/${host.id}` },
                position: { x: 500, y: index * 200 },
                style: {
                    backgroundColor: useLayout().Colors.hostColor,
                    color: "#000",
                    width: `${width + 20}px`, // Adding some padding
                    height: `${height + 20}px`, // Adding some padding
                    border: `3px solid ${borderColor}`
                },
                type: hostNodeType,
                sourcePosition: 'right',
                targetPosition: 'left',
            };
        });

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
