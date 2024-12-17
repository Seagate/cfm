// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import dagre from '@dagrejs/dagre';
import { Position } from '@vue-flow/core';
import { ref } from 'vue';

/**
 * Composable to run the layout algorithm on the graph.
 * It uses the `dagre` library to calculate the layout of the nodes and edges.
 */

export function measureText(text: string, font = '16px Arial') {
  const canvas = document.createElement('canvas');
  const context = canvas.getContext('2d');
  context!.font = font;
  const metrics = context!.measureText(text);
  const width = Math.max(metrics.width, 220);
  return {
    width: width,
    height: parseInt(font, 10) // Assuming height is roughly the font size
  };
}

export const Colors = {
  applianceColor: '#f2ae72',
  baldeColor: '#f2e394',
  hostColor: '#d9ecd0',
  serviceColor: '#6ebe4a',
};

export function borderColorChange(status: string | undefined) {
  switch (status) {
    case "online":
      return "#6ebe4a"; // Green
    case "offline":
      return "#b00020"; // Red
    case "unavailable":
      return "#ff9f40"; // Orange
    default:
      return "#B0B0B0"; // Gray
  }
}

export function useLayout() {
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
      // Measure the text dimensions for dynamic sizing
      const { width, height } = measureText(node.data.label);

      dagreGraph.setNode(node.id, {
        width: width + 20, // Adding some padding
        height: height + 20 // Adding some padding
      });
    }

    for (const edge of edges) {
      dagreGraph.setEdge(edge.source, edge.target);
    }

    dagre.layout(dagreGraph);

    // Set nodes with updated positions
    return nodes.map((node: { id: string | dagre.Label; style: any; }) => {
      const nodeWithPosition = dagreGraph.node(node.id);

      return {
        ...node,
        targetPosition: isHorizontal ? Position.Left : Position.Top,
        sourcePosition: isHorizontal ? Position.Right : Position.Bottom,
        position: { x: nodeWithPosition.x, y: nodeWithPosition.y },
        style: {
          ...node.style,
          width: `${nodeWithPosition.width}px`,
          height: `${nodeWithPosition.height}px`
        }
      };
    });
  }

  return { graph, layout, previousDirection, measureText, borderColorChange, Colors };
}