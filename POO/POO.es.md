# POO


## Pilares de la programación orientada a objetos
### Herencia
La herencia es un mecanismo de la Programación Orientada a Objetos (POO) que permite que una clase (subclase) herede atributos y métodos de otra clase (superclase), promoviendo la reutilización de código.
```c#

class Animal {
    public virtual void hacerSonido() {
        Console.WriteLine("Sonido de animal");
    }
}

class Perro: Animal {
    public void ladrar() {
        Console.WriteLine("Guau Guau");
    }
}


class Program {
    public static void Main() {
        Perro perro = new Perro();
        perro.hacerSonido(); // Salida: Sonido de animal
        perro.ladrar(); // Salida: Guau Guau
    }
}

```

### Polimorfismo 
El polimorfismo permite que un mismo método tenga diferentes implementaciones dependiendo de la clase que lo utiliza.

```c#

class Animal {
    public virtual void hacerSonido() {
        Console.WriteLine("Sonido de animal");
    }
}

class Gato: Animal {
    public override void hacerSonido() {
        Console.WriteLine("Miau");
    }
}


class Program {
    public static void Main() {
        Animal animal = new Gato();
        animal.hacerSonido(); // Salida: Miau
    }
}

```

### Encapsulamiento
El encapsulamiento es el principio de la POO que protege los datos internos de un objeto, permitiendo acceso controlado a través de métodos públicos y propiedades.

```c#
class CuentaBancaria {
    private double saldo;
    
    public CuentaBancaria(double saldoInicial) {
        saldo = saldoInicial;
    }
    
    public void Depositar(double valor) {
        if (valor > 0) {
            saldo += valor;
            Console.WriteLine("Depósito realizado con éxito.");
        }
    }
    
    public void Retirar(double valor) {
        if (valor > 0 && valor <= saldo) {
            saldo -= valor;
            Console.WriteLine("Retiro realizado con éxito.");
        } else {
            Console.WriteLine("Saldo insuficiente.");
        }
    }
}

class Program {
    public static void Main() {
        CuentaBancaria cuenta = new CuentaBancaria(1000);
        cuenta.Depositar(500); // Salida: Depósito realizado con éxito.
        cuenta.Retirar(200); // Salida: Retiro realizado con éxito.
        cuenta.Retirar(2000); // Salida: Saldo insuficiente.
    }
}
```
