# POO


## Pilares da orientação a objetos
### Herança
Herança é um mecanismo da Programação Orientada a Objetos (POO) que permite que uma classe (subclasse) herde atributos e métodos de outra classe (superclasse), promovendo a reutilização de código.
```c#

class Animal {
    public virtual void fazerSom() {
        Console.WriteLine("Som de animal");
    }
}

class Cachorro: Animal {
    public void latir() {
        Console.WriteLine("Au Au");
    }
}


class Program {
    public static void Main() {
        Cachorro cachorro = new Cachorro();
        cachorro.fazerSom(); // Saída: Som de animal
        cachorro.latir(); // Saída: Au Au
    }
}

```

### Polirmorfismo 
Polimorfismo permite que um mesmo método tenha diferentes implementações dependendo da classe que o utiliza.

```c#

class Animal {
    public virtual void fazerSom() {
        Console.WriteLine("Som de animal");
    }
}

class Gato: Animal {
    public override void emitirSom() {
        Console.WriteLine("Miau");
    }
}


class Program {
    public static void Main() {
        Animal animal = new Gato();
        animal.fazerSom(); // Saída: Miau
    }
}

```

### Encapsulamento
Encapsulamento é o princípio da POO que protege os dados internos de um objeto, permitindo acesso controlado através de métodos públicos e propriedades.

```
class ContaBancária {
    private double saldo;
    
    public ContaBancária(double saldoInicial) {
        saldo = saldoInicial;
    }
    
    public void Depositar(double valor) {
        if (valor > 0) {
            saldo += valor;
            Console.WriteLine("Depósito realizado com sucesso.");
        }
    }
    
    public void Sacar(double valor) {
        if (valor > 0 && valor <= saldo) {
            saldo -= valor;
            Console.WriteLine("Saque realizado com sucesso.");
        } else {
            Console.WriteLine("Saldo insuficiente.");
        }
    }

}

class Program {
    public static void Main() {
        ContaBancária conta = new ContaBancária(1000);
        conta.Depositar(500); // Saída: Depósito realizado com sucesso.
        conta.Sacar(200); // Saída: Saque realizado com sucesso.
        conta.Sacar(2000); // Saída: Saldo insuficiente.
    }
}
```