import algoritms.LinkedList;

class HelloWorld {
  public static void main(String[] args) {
    LinkedList<Integer> list = new LinkedList<>();
    list.append(10);
    list.unshift(5);
    list.appendOnIndex(7, 1);

    list.printList(); // 5 -> 7 -> 10 

   var indexFound = list.search(null);
    if (indexFound != null) {
      System.out.println("Value found at index: " + indexFound);
    } else {
      System.out.println("Value not found");
    }

   list.delete(7);
    list.printList(); // 5 -> 10

  }
}