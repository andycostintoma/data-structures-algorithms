def main():
    # Initialize the array
    arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
    print("After initialization:", arr)

    # O(1) - Access element at index 5
    try:
        value = arr[5]  # Access the value at index 5
        print("Accessed value at index 5:", value)
    except IndexError as e:
        print("Access error:", e)

    # O(n) - Search for the value 1
    try:
        found_index = arr.index(1)  # Find the index of value 1
        print(f"Found 1 at index: {found_index}")
    except ValueError:
        found_index = -1  # Value not found
        print("1 not found in the array")

    # O(n) - Insert 99 at index 5
    try:
        arr.insert(5, 99)  # Insert 99 at index 5
        print("After insert at index 5:", arr)
    except IndexError as e:
        print("Insert error:", e)

    # O(n) - Delete element at index 2
    try:
        del arr[2]  # Delete the element at index 2
        print("After delete at index 2:", arr)
    except IndexError as e:
        print("Delete error:", e)

    # O(1) - Append 100 to the end of the array
    arr.append(100)
    print("After appending 100:", arr)

    # O(1) - Pop the last element
    try:
        popped = arr.pop()  # Pop the last element from the array
        print(f"Popped value: {popped}")
        print("After popping:", arr)
    except IndexError as e:
        print("Pop error:", e)

    # Final array
    print("Final array:", arr)


if __name__ == "__main__":
    main()
