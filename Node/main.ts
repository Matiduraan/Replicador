import { MOCK_FUNCTIONS } from "./mock_functions";
import tree from "./tree.json";

type Node = {
  id: string;
  name: string;
  params: string[];
  children: string[]; // IDs of the children
};

const MOCK_STEPS: string[] = ["1", "5", "3", "4", "2"];

/* 
  El arbol sera leido de forma recursiva, generando una estructura de datos donde para cada nodo dentro de los mock steps, se indiquen los parametros, el path y de que nodos depende
  Para esto se creara una funcion que reciba un nodo y devuelva un objeto con la siguiente estructura:
  {
    id: string,
    name: string,
    params: string[],
    dependencies: string[] -> IDs de los nodos que dependen de el
  }
  
*/

function verifyDependencies(nodesToExecute: ProcessedNode[]): boolean {
  return nodesToExecute.every((node) =>
    node.dependencies.every((dependency) =>
      nodesToExecute.some((finalNode) => finalNode.id === dependency)
    )
  );
}

function buildDependencyMap(tree: Node[]): Record<string, string[]> {
  const dependencyMap: Record<string, string[]> = {};

  tree.forEach((node) => {
    node.children.forEach((childId) => {
      if (!dependencyMap[childId]) {
        dependencyMap[childId] = [];
      }
      dependencyMap[childId].push(node.id);
    });
  });

  return dependencyMap;
}

type ProcessedNode = {
  id: string;
  name: string;
  params: string[];
  dependencies: string[];
};

function processNode(
  node: Node,
  dependencyMap: Record<string, string[]>
): Node & { dependencies: string[] } {
  const { id } = node;
  const dependencies = dependencyMap[id] || [];
  return { ...node, dependencies };
}

function processTree(tree: Node[], stepsToExecute: string[]): ProcessedNode[] {
  const dependencyMap = buildDependencyMap(tree);
  return tree
    .map((node) => processNode(node, dependencyMap))
    .filter((node) =>
      stepsToExecute.some((mock_step) => mock_step === node.id)
    );
}

function findNodesToExecute(
  finishedNodes: ProcessedNode[],
  processedNodes: ProcessedNode[]
): ProcessedNode[] {
  const finishedNodeIds = new Set(finishedNodes.map((node) => node.id));

  return processedNodes.filter((node) => {
    return (
      !finishedNodeIds.has(node.id) &&
      node.dependencies.every((dependency) => finishedNodeIds.has(dependency))
    );
  });
}

function executionLoop() {
  const processedNodes = processTree(tree.nodes, MOCK_STEPS);
  // console.log("Processed nodes", processedNodes);
  if (!verifyDependencies(processedNodes)) {
    throw new Error("There are circular or missing dependencies in the tree");
  }
  const finishedNodes: ProcessedNode[] = [];
  let paramsHeap: Record<string, unknown> = {};
  let n = 0;
  while (finishedNodes.length < processedNodes.length) {
    const nodesToExecute = findNodesToExecute(finishedNodes, processedNodes);

    nodesToExecute.forEach((node) => {
      // Execute the node
      const functionParams = {};
      node.params.forEach((param) => {
        if (!paramsHeap[param]) {
          throw new Error(`Missing parameter ${param} for node ${node.name}`);
        }
        functionParams[param] = paramsHeap[param];
      });

      const result = MOCK_FUNCTIONS[node.name]({ ...functionParams });
      paramsHeap = { ...paramsHeap, ...result };
      finishedNodes.push(node);
    });
    n++;
  }
}

executionLoop();
// const result = processTree(tree.nodes, MOCK_STEPS);

// console.log(result);
