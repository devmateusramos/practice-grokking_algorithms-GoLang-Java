package factorial;

import org.junit.Assert;
import org.junit.Test;
// Junit4

public class Factorial2Test {
    @Test
    public void testFactorial() {
        Factorial2 factorial2 = new Factorial2();

        Assert.assertEquals(120, factorial2.getFactorial(5));
    }
}
