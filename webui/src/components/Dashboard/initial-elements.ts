// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { computed } from "vue";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useServiceStore } from "../Stores/ServiceStore";
import { useLayout } from "./useLayout";

export const useFlowData = () => {
  const applianceStore = useApplianceStore();
  const hostStore = useHostStore();
  const serviceStore = useServiceStore();

  const position = { x: 0, y: 0 };
  const serviceNodeType = "cfm-service";
  const applianceNodeType = "appliance";
  const hostNodeType = "host";
  const bladeNodeType = "blade";

  const nodes = computed(() => {
    const coreNode = serviceStore.serviceVersion
      ? [
        {
          id: "cfm-service",
          data: { label: "CFM Service", },
          position: position,
          style: { backgroundColor: "#6ebe4a", color: "#000" },
          type: serviceNodeType,
        },
      ]
      : [];

    const applianceNodes = applianceStore.applianceIds.flatMap(
      (appliance, applianceIndex) => [
        {
          id: `appliance-${appliance.id}`,
          data: { label: appliance.id, url: `/appliances/${appliance.id}` },
          position: position,
          style: { backgroundColor: "#f2ae72", color: "#000" },
          type: applianceNodeType,
        },
        ...appliance.bladeIds.map((bladeId, bladeIndex) => ({
          id: `blade-${bladeId}`,
          data: { label: bladeId, url: `/appliances/${appliance.id}/blades/${bladeId}`, associatedAppliance: appliance.id },
          position: position,
          style: { backgroundColor: "#f2e394", color: "#000" },
          type: bladeNodeType,
        })),
      ]
    );

    const hostNodes = hostStore.hostIds.map((host, index) => ({
      id: `host-${host}`,
      data: { label: host, url: `/hosts/${host}` },
      position: position,
      style: { backgroundColor: "#d9ecd0", color: "#000" },
      type: hostNodeType,
    }));

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

    // Apply the layout
    return useLayout().layout(allNodes, edges.value, 'LR');
  });

  const edges = computed(() => {
    const coreEdges = serviceStore.serviceVersion
      ? applianceStore.applianceIds.flatMap((appliance) => [
        {
          id: `cfm-appliance-${appliance.id}`,
          source: "cfm-service",
          target: `appliance-${appliance.id}`,
          animated: true,
        },
        ...appliance.bladeIds.map((bladeId) => ({
          id: `appliance-blade-${appliance.id}-${bladeId}`,
          source: `appliance-${appliance.id}`,
          target: `blade-${bladeId}`,
          animated: true,
        })),
      ])
      : [];

    const hostEdges = hostStore.hostIds.map((host) => ({
      id: `cfm-${host}`,
      source: "cfm-service",
      target: `host-${host}`,
      animated: true,
    }));

    return [...coreEdges, ...hostEdges];
  });

  return { nodes, edges };
};
