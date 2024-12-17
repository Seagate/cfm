// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { computed } from "vue";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useServiceStore } from "../Stores/ServiceStore";
import { useLayout } from "./useLayout";

export const useControlData = () => {
  const applianceStore = useApplianceStore();
  const hostStore = useHostStore();
  const serviceStore = useServiceStore();

  const position = { x: 0, y: 0 };
  const serviceNodeType = "cfm-service";
  const applianceNodeType = "appliance";
  const hostNodeType = "host";
  const bladeNodeType = "blade";

  const controlNodes = computed(() => {
    const coreNode = serviceStore.serviceVersion
      ? [
        {
          id: "cfm-service",
          data: { label: "CFM Service", },
          position: position,
          style: { backgroundColor: useLayout().Colors.serviceColor, border: "none" },
          type: serviceNodeType,
        },
      ]
      : [];

    const applianceNodes = applianceStore.applianceIds.flatMap(
      (appliance) => [
        {
          id: `appliance-${appliance.id}`,
          data: { label: appliance.id, url: `/appliances/${appliance.id}` },
          position: position,
          style: { backgroundColor: useLayout().Colors.applianceColor, border: "none" },
          type: applianceNodeType,
        },
        ...appliance.blades.map((blade) => {
          const borderColor = useLayout().borderColorChange(blade.status);
          return {
            id: `blade-${blade.id}`,
            data: { label: blade.id, url: `/appliances/${appliance.id}/blades/${blade.id}`, associatedAppliance: appliance.id },
            position: position,
            style: { backgroundColor: useLayout().Colors.baldeColor, border: `3px solid ${borderColor}` },
            type: bladeNodeType,
          }
        }),
      ]
    );

    const hostNodes = hostStore.hostIds.map((host) => {
      const borderColor = useLayout().borderColorChange(host.status);

      return {
        id: `host-${host.id}`,
        data: { label: host.id, url: `/hosts/${host.id}` },
        position: position,
        style: { backgroundColor: useLayout().Colors.hostColor, border: `3px solid ${borderColor}` },
        type: hostNodeType,
      }
    });

    const allNodes = [...coreNode, ...applianceNodes, ...hostNodes];

    const edges = computed(() => {
      const coreEdges = serviceStore.serviceVersion
        ? applianceStore.applianceIds.flatMap((appliance) => [
          {
            id: `cfm-appliance-${appliance.id}`,
            type: "smoothstep",
            source: "cfm-service",
            target: `appliance-${appliance.id}`,
          },
          ...appliance.blades.map((blade) => ({
            id: `appliance-blade-${appliance.id}-${blade.id}`,
            source: `appliance-${appliance.id}`,
            target: `blade-${blade.id}`,
          })),
        ])
        : [];

      const hostEdges = hostStore.hostIds.map((host) => ({
        id: `cfm-${host.id}`,
        source: "cfm-service",
        target: `host-${host.id}`,
      }));

      return [...coreEdges, ...hostEdges];
    });

    // Apply the layout
    return useLayout().layout(allNodes, edges.value, 'LR');
  });

  const controlEdges = computed(() => {
    const coreEdges = serviceStore.serviceVersion
      ? applianceStore.applianceIds.flatMap((appliance) => [
        {
          id: `cfm-appliance-${appliance.id}`,
          source: "cfm-service",
          target: `appliance-${appliance.id}`,
          animated: true,
        },
        ...appliance.blades.map((blade) => ({
          id: `appliance-blade-${appliance.id}-${blade.id}`,
          source: `appliance-${appliance.id}`,
          target: `blade-${blade.id}`,
          animated: true,
        })),
      ])
      : [];

    const hostEdges = hostStore.hostIds.map((host) => ({
      id: `cfm-${host.id}`,
      source: "cfm-service",
      target: `host-${host.id}`,
      animated: true,
    }));

    return [...coreEdges, ...hostEdges];
  });

  return { controlNodes, controlEdges };
};
