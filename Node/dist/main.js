"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const mock_functions_1 = require("./mock_functions");
const tree_json_1 = __importDefault(require("./tree.json"));
const MOCK_STEPS = ["1", "5", "3", "4", "2"];
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
function verifyDependencies(nodesToExecute) {
    return nodesToExecute.every((node) => node.dependencies.every((dependency) => nodesToExecute.some((finalNode) => finalNode.id === dependency)));
}
function buildDependencyMap(tree) {
    const dependencyMap = {};
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
function processNode(node, dependencyMap) {
    const { id } = node;
    const dependencies = dependencyMap[id] || [];
    return Object.assign(Object.assign({}, node), { dependencies });
}
function processTree(tree, stepsToExecute) {
    const dependencyMap = buildDependencyMap(tree);
    console.log("DEPENDENCY_MAP", dependencyMap);
    return tree
        .map((node) => processNode(node, dependencyMap))
        .filter((node) => stepsToExecute.some((mock_step) => mock_step === node.id));
}
function findNodesToExecute(finishedNodes, processedNodes) {
    const finishedNodeIds = new Set(finishedNodes.map((node) => node.id));
    return processedNodes.filter((node) => {
        return (!finishedNodeIds.has(node.id) &&
            node.dependencies.every((dependency) => finishedNodeIds.has(dependency)));
    });
}
function executionLoop() {
    const processedNodes = processTree(tree_json_1.default.nodes, MOCK_STEPS);
    // console.log("Processed nodes", processedNodes);
    if (!verifyDependencies(processedNodes)) {
        throw new Error("There are circular or missing dependencies in the tree");
    }
    const finishedNodes = [];
    let paramsHeap = {};
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
            const result = mock_functions_1.MOCK_FUNCTIONS[node.name](Object.assign({}, functionParams));
            paramsHeap = Object.assign(Object.assign({}, paramsHeap), result);
            finishedNodes.push(node);
        });
        n++;
    }
}
executionLoop();
// const result = processTree(tree.nodes, MOCK_STEPS);
// console.log(result);
//# sourceMappingURL=main.js.map