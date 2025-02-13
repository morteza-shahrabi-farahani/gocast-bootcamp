\* One of the benefits of using enums as type is that we can define methods and functions for them. They also increase code readability.

\* using if initializations for error variables is better than defining them. Because beyond the scope of the if statement, we do not need the error variable anymore.

\* For wrapping errors of the lower layers and also adding your message in this layer to the error, you can use, ``` fmt.Errorf("Unexpected error %w", err)```