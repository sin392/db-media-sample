import re
import sys

COLORS = [
    # adapter
    'darkgreen',
    # domain
    'darkgoldenrod',
    # infrastructure
    'darkblue',
    # usecase
    'darkred',
    # others
    'darkorange',
    'mediumvioletred',
    'dimgray',
    'ind',
    'darkcyan',
    'black'
]


def color_nodes(graph: str) -> str:
    color_line_map = {}
    current_color = 0
    results = []
    for raw_line in graph.split('\n'):
        line = raw_line.strip()
        # ノードに関する行のみを処理
        if line.startswith('"') and '->' not in line:
            node_name = line.split(' ')[0].replace('"', '')
            # ノードのprefixを取得
            prefix = node_name.split('/')[0]
            # prefixに対して色を割り当てる
            if prefix not in color_line_map:
                color = COLORS[current_color]
                color_line_map[prefix] = color
                current_color += 1
            # すでに色が割り当てられている場合はそのまま追加
            else:
                color = color_line_map[prefix]
            # 色を上書き
            color_line = re.sub(
                r'fillcolor="[0-9\.]+\s[0-9\.]+\s[0-9\.]+"',
                f'fillcolor="{color}"',
                raw_line
            )
            results.append(color_line)
        # ノード以外の行はそのまま追加
        else:
            results.append(raw_line)
    return '\n'.join(results)


if __name__ == '__main__':
    graph = sys.stdin.read()
    colored_graph = color_nodes(graph)
    print(colored_graph)
