import java.util.*;

public class UnionFind {

    // --- DSU (Disjoint Set Union) with path compression and union by rank ---
    static int[] parent, rank;

    static void init(int n) {
        parent = new int[n];
        rank = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
    }

    static int find(int x) {
        if (parent[x] != x) parent[x] = find(parent[x]);
        return parent[x];
    }

    static boolean union(int x, int y) {
        int px = find(x), py = find(y);
        if (px == py) return false;
        if (rank[px] < rank[py]) { int t = px; px = py; py = t; }
        parent[py] = px;
        if (rank[px] == rank[py]) rank[px]++;
        return true;
    }

    // 1. KruskalMST — minimum spanning tree total weight (CLRS Chapter 21)
    public static int kruskalMST(int n, int[][] edges) {
        Arrays.sort(edges, (a, b) -> a[2] - b[2]);
        init(n);
        int total = 0;
        for (int[] e : edges) {
            if (union(e[0], e[1])) total += e[2];
        }
        return total;
    }

    // 2. CountComponents — count connected components (LeetCode #323)
    public static int countComponents(int n, int[][] edges) {
        init(n);
        for (int[] e : edges) union(e[0], e[1]);
        Set<Integer> roots = new HashSet<>();
        for (int i = 0; i < n; i++) roots.add(find(i));
        return roots.size();
    }

    // 3. FindRedundantConnection — return the edge forming a cycle (LeetCode #684)
    public static int[] findRedundantConnection(int[][] edges) {
        init(edges.length + 1);
        for (int[] e : edges) {
            if (!union(e[0], e[1])) return e;
        }
        return new int[0];
    }

    // 4. AccountsMerge — merge accounts sharing an email (LeetCode #721)
    public static List<List<String>> accountsMerge(List<List<String>> accounts) {
        Map<String, Integer> emailToId = new HashMap<>();
        Map<String, String> emailToName = new HashMap<>();
        int id = 0;
        for (List<String> acc : accounts) {
            String name = acc.get(0);
            for (int i = 1; i < acc.size(); i++) {
                String email = acc.get(i);
                if (!emailToId.containsKey(email)) emailToId.put(email, id++);
                emailToName.put(email, name);
            }
        }
        init(id);
        for (List<String> acc : accounts) {
            int first = emailToId.get(acc.get(1));
            for (int i = 2; i < acc.size(); i++) {
                union(first, emailToId.get(acc.get(i)));
            }
        }
        // Group by root
        Map<Integer, List<String>> groups = new HashMap<>();
        for (String email : emailToId.keySet()) {
            int root = find(emailToId.get(email));
            groups.computeIfAbsent(root, k -> new ArrayList<>()).add(email);
        }
        List<List<String>> result = new ArrayList<>();
        for (List<String> emails : groups.values()) {
            Collections.sort(emails);
            List<String> merged = new ArrayList<>();
            merged.add(emailToName.get(emails.get(0)));
            merged.addAll(emails);
            result.add(merged);
        }
        return result;
    }
}
