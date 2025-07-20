import java.io.*;
import java.util.*;


public class Main{
    final static int MAX_N = (int)1e5+5;
    static int n,m,d,p;
    static boolean affected[] = new boolean [MAX_N];
    static ArrayList <Integer> graph[] = new ArrayList [MAX_N];
    static int maxDown[] = new int[MAX_N];
    static int maxUp[] = new int[MAX_N];



    public static void main(String[] args) {
        Read();
        System.out.println(Solve());
    }

    static int Solve(){
        ComputeMaxDown(0,-1);
        maxUp[0] = affected[0] == true ? 0 : -1;
        ComputeMaxUp(0,-1);
        int ans = 0;
        for(int i = 0 ; i < n ; i++){
            if(maxDown[i] <= d && maxUp[i] <= d)
                ans++;
        }
        return ans;
    }

    static void ComputeMaxUp(int src, int par){
        int max1 = -1, max2 = -1;
        for(int i = 0 ; i < graph[src].size(); i++){
            int v = graph[src].get(i);
            if(v == par) continue;
            if(maxDown[v] > max1){ max2 = max1; max1 = maxDown[v];}
            else if(maxDown[v] > max2) {max2 = maxDown[v];}
        }

        for(int i = 0 ; i < graph[src].size(); i++){
            int v = graph[src].get(i);
            if(v == par) continue;
            maxUp[v] = (maxDown[v] == max1 ? max2 : max1);
            if(maxUp[v] != -1 )// in case of root node or  sibling node have no affective node
                    maxUp[v] += 2;
            if(maxUp[src] != -1) // in case of root node...
                    maxUp[v] = Math.max(maxUp[v],maxUp[src] +1);
            if(affected[v])
                    maxUp[v] = Math.max(maxUp[v],0);
            ComputeMaxUp(v,src);
        }
    }

    static int ComputeMaxDown(int src, int par){
        if(affected[src] == false){
            maxDown[src] = -1;
        }
        else maxDown[src] = 0;

        for(int i = 0 ; i < graph[src].size(); i++){
            int v = graph[src].get(i);
            if(v==par) continue;
            int d = ComputeMaxDown(v,src);
            if(d > -1) // -1 is the assign that there are not such a affected node in the subtree of src
            maxDown[src] = Math.max(maxDown[src],d+1);
        }
        return maxDown[src];
    }

    static void Read(){
        Scanner sc = new Scanner(System.in);
        int u,v;
        n = sc.nextInt();
        m = sc.nextInt();
        d = sc.nextInt();
        for(int i = 0 ; i < n ; i++){
            graph[i] = new ArrayList<>();
        }
        for(int i = 0 ; i < m ; i++){
            p = sc.nextInt();
            affected[--p] = true;
        }
        for(int i = 0 ; i < n-1; i++){
            u = sc.nextInt();
            v = sc.nextInt();
            --u;--v;
            graph[u].add(v);
            graph[v].add(u);
        }
    }
}
