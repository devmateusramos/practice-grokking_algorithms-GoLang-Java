package factorial;

public class Factorial2 {
    public static void main(String[] args) {
        Factorial2 factorial = new Factorial2();
        System.out.println("Fatorial de 5 é: " + factorial.getFactorial(5));
    }

    public int getFactorial(int number) {
        if(number == 1 || number == 0) {
            return 1;
        }

        return number * getFactorial(number - 1);
    }
}
