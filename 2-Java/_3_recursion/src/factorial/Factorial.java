package factorial;

public class Factorial {
    public static void main(String[] args) {
        System.out.println(fact(0));
    }
    private static int fact(int x) {
        if (x == 1 || x == 0) {
            return 1;
        }

        return x * fact(x - 1);
    }
}
