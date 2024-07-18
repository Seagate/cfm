// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import dagre from '@dagrejs/dagre';
import { Position, useVueFlow } from '@vue-flow/core';
import { ref } from 'vue';

/**
 * Composable to run the layout algorithm on the graph.
 * It uses the `dagre` library to calculate the layout of the nodes and edges.
 */
export function useLayout() {
  const { findNode } = useVueFlow();

  const graph = ref(new dagre.graphlib.Graph());

  const previousDirection = ref('LR');

  function layout(nodes: any[], edges: any, direction: string) {
    // Create a new graph instance
    const dagreGraph = new dagre.graphlib.Graph();

    graph.value = dagreGraph;

    dagreGraph.setDefaultEdgeLabel(() => ({}));

    const isHorizontal = direction === 'LR';
    dagreGraph.setGraph({ rankdir: direction });

    previousDirection.value = direction;

    for (const node of nodes) {
      // Use the dimensions property of the internal node (`GraphNode` type)
      const graphNode = findNode(node.id);

      dagreGraph.setNode(node.id, { width: graphNode?.dimensions.width || 150, height: graphNode?.dimensions.height || 50 });
    }

    for (const edge of edges) {
      dagreGraph.setEdge(edge.source, edge.target);
    }

    dagre.layout(dagreGraph);

    // Set nodes with updated positions
    return nodes.map((node: { id: string | dagre.Label; }) => {
      const nodeWithPosition = dagreGraph.node(node.id);

      return {
        ...node,
        targetPosition: isHorizontal ? Position.Left : Position.Top,
        sourcePosition: isHorizontal ? Position.Right : Position.Bottom,
        position: { x: nodeWithPosition.x, y: nodeWithPosition.y },
      };
    });
  }

  return { graph, layout, previousDirection };
}
