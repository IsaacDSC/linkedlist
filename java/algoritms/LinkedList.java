package algoritms;

import java.util.Optional;

class Node<T> {
  T data;
  Node<T> next;

  Node(T data) {
    this.data = data;
    this.next = null;
  }
}

public class LinkedList<T> {
  private Node<T> head;

  public LinkedList() {
    this.head = null;
  }

  public void unshift(T data) {
    Node<T> newNode = new Node<>(data);
    if (head == null) {
      head = newNode;
    } else {
      newNode.next = head;
      head = newNode;
    }
  }

  public void append(T data) {
    Node<T> newNode = new Node<>(data);
    if (head == null) {
      head = newNode;
    } else {
      Node<T> current = head;
      while (current.next != null) {
        current = current.next;
      }
      current.next = newNode;
    }
  }

  public void appendOnIndex(T value, int index) {
    Node<T> newNode = new Node<>(value);
    if (index == 0) {
      unshift(value);
      return;
    }
    Node<T> current = head;
    for (int i = 0; i < index - 1 && current != null; i++) {
      current = current.next;
    }
    if (current != null) {
      newNode.next = current.next;
      current.next = newNode;
    } else {
      System.out.println("Index out of bounds");
    }
  }

  public Optional<Integer> search(T value) {
    Node<T> current = head;
    var index = 0;
    while (current != null) {
      if (current.data.equals(value)) {
        System.out.println("Found: " + value);
        return Optional.of(index);
      }
      current = current.next;
      index++;
    }
    return null;
  }

  public void delete(T value) {
    if (head == null) {
      System.out.println("List is empty");
      return;
    }
    if (head.data.equals(value)) {
      head = head.next;
      return;
    }
    Node<T> current = head;
    while (current.next != null && !current.next.data.equals(value)) {
      current = current.next;
    }
    if (current.next != null) {
      current.next = current.next.next;
    } else {
      System.out.println("Value not found: " + value);
    }
  }

  public void deleteOnIndex(int index) {
    if (head == null) {
      System.out.println("List is empty");
      return;
    }
    if (index == 0) {
      head = head.next;
      return;
    }
    Node<T> current = head;
    for (int i = 0; i < index - 1 && current != null; i++) {
      current = current.next;
    }
    if (current != null && current.next != null) {
      current.next = current.next.next;
    } else {
      System.out.println("Index out of bounds");
    }
  }

  public void reverse() {
    Node<T> prev = null;
    Node<T> current = head;
    Node<T> next = null;
    while (current != null) {
      next = current.next;
      current.next = prev;
      prev = current;
      current = next;
    }
    head = prev;
  }

  public void printList() {
    Node<T> current = head;
    while (current != null) {
      System.out.print(current.data + " ");
      current = current.next;
    }
    System.out.println();
  }

}
