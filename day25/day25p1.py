import re
import networkx as nx
import math

parse = lambda line: re.findall(r'\w+', line)

def make_graph(INPUT: list[str]) -> nx.Graph:
    G = nx.Graph()
    for key, *vals in (parse(line) for line in INPUT):
        G.add_edges_from([(key, val) for val in vals])
    return G

def solve(INPUT: list[str], show: bool = False) -> None:
    G = make_graph(INPUT)
    if show:
        nx.draw_networkx(G, pos=nx.spring_layout(G))

    G.remove_edges_from(nx.minimum_edge_cut(G))
    print(math.prod(map(len, nx.connected_components(G))))

solve(open(0))
