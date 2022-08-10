package _1_Countdown;

public class Countdown {
    public static void main(String[] args) {
        countdown(5);
    }
    private static void countdown(int i) {
        System.out.println(i);

        if (i <= 0 ) {
            return;
        } else {
            countdown(i - 1);
        }
    }
}
