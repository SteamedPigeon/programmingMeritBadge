// Example F to C in Java
// This file must be named FahrenheitToCelsius.java 
import java.util.Scanner;
public class FahrenheitToCelsius {
    public static final  double  LOW_TEMP_F_WARNING=0.;
    public static final  double  HIGH_TEMP_F_WARNING=100.;
    public static final  int     MAX_LOOP=5;
    public static void main(String[] args) {
        Scanner scanFaren = new Scanner(System.in);
        double Farenheit = 0;
        double Celsius = 0;
        for(int i=0; i<MAX_LOOP; i++){
        System.out.print("\nEnter a temperature in Fahrenheit: ");
            if(scanFaren.hasNextDouble())
                {
                    Farenheit=scanFaren.nextDouble();
                    Celsius = ( Farenheit- 32.)*5./9.;
                }else{
                System.out.println("Data entry error - try again\n");
                System.exit(-1); }
            System.out.println("The temperature in Celsius is: "+Celsius);
       
            // Check for high temperature and issue a warning if necessary
            if(Farenheit > HIGH_TEMP_F_WARNING){
                System.out.print("Remember to hydrate\n");}
            // Check for low temperature and issue a warning if necessary
            if(Farenheit < LOW_TEMP_F_WARNING ){
                System.out.print("Remember to pack Long underwear\n");}
        }
        scanFaren.close();
        System.exit(-1);
    }
}