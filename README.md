# MessagePack vs JSON

MessagePack and JSON store data similarly. Any JSON blob can be encoded in MessagePack. The single advantage of JSON is that it’s human-readable, though that’s also a disadvantage because people sometimes try to construct JSON by hand and get it wrong. You should use the MessagePack format if you’re interested in efficiency, the size of the encoded data, and additional features such as encoding raw binary data efficiently.

Take numbers, for example. The number 5 takes one byte in both formats, just as all the numbers between 0 and 9. The number 10 takes two bytes in JSON (it’s two characters) but only one byte in MessagePack. The number 100 requires three bytes in JSON but still only one byte in MessagePack. The number -108 takes four bytes in JSON but one byte in MessagePack. So JSON simply requires as many bytes as there are characters in the number in decimal format, including the dot for non-integers and the minus sign for numbers below zero. On the other hand, all integers between -32 and 127 take only one byte in MessagePack. Big numbers (and very negative numbers) take up can require much more space in the JSON format.

Computers represent numbers not as characters (like '0', '1', '2') but in a binary format, and that’s how MessagePack does it. A single byte (8 bits) can represent a signed integer and thus hold any integer between -128 and 127, or an unsigned integer and thus represent any integer between 0 and 255. Why then can’t MessagePack just use one byte for a number like -50 or 190, using the appropriate data type to encode the number? The reason is that MessagePack data is self-contained (or self-describing) and encodes the data types directly within the message—so it does not require an external representation of the types of data contained in the encoded blob. (This is a great feature that distinguishes MessagePack from the Protocol Buffers format, though it does mean that MessagePack can require slightly more space than Protocol Buffers.)

What about strings? Since JSON data is all strings, you might think it’s trivial to represent a string in JSON: simply wrap the string in double quotes. If you’ve ever constructed JSON messages this way, you’ve probably run into strange errors. The problem is that there are certain “text” characters that need to be treated specifically in a JSON string. One obvious character is the double quote ("), which wraps the string; all occurrences of the character within the string must be replaced by a backslash and a double quote (\\"). This is called character escaping. Other characters that need escaping include newlines and tabs, and there are others. So a string that takes only 15 bytes in binary may take twice as much, 30 bytes, in JSON. How? If it consists entirely of characters that need to be escaped. So your message size can grow significantly with JSON, and it always takes at least two extra bytes for the quote characters that wrap the string. And you need to take special care when encoding and decoding certain characters, which slows the process down.

But you don’t have this problem with binary encoding formats like MessagePack: there’s no escaping necessary. A string in MessagePack simply requires between one and nine extra bytes to indicate that you’ve got a string-typed element and the length of the string in bytes.

What about binary data, such as images and sound files and everything else that is not text? To encode binary data in the JSON format, you first need to encode it in a text format that perfectly preserves all the data. Typically base-64 or base-16 (hexadecimal) encoding is used to convert the binary data to a string, and a JSON string is encoded. The encoded string *always* takes up more space than the raw binary array: base-16 requires twice as many bytes, and base-64 requires four bytes for every three bytes of the raw data. This means extra overhead upon encoding, transferring, and decoding the data and is not recommended for performance-sensitive applications.

On the other hand, the binary data format fits the MessagePack format perfectly, as it’s encoded pretty much the same way as strings. There’s no extra converting, no checking for characters to escape.

What about arrays and key-value pairs (maps)? Think about the extra characters required with JSON. A comma must separate all elements, and a colon separates a map key from its value. And a pair of square brackets or curly braces ([] or {}) are needed to enclose the array or map. Again, this means extra storage space requirements and more data to transfer. With MessagePack, you only specify the data type (array/map) along with the number of elements present. This can be *just one* extra byte in all for a small arrays and maps.

There’s more to say on this topic. I suggest you try using MessagePack. There are great libraries in different programming languages that encode and decode data. If you’ll allow me to plug mine, check out: [github.com/dchenk/msgp](https://github.com/dchenk/msgp)

## Size of Encoded Messages

Let's consider the size of encoded messages in each of the two formats. We want to compare small, medium, and large messages to see which format is best for which kinds of messages. (The messages for the JSON size calculation are minified.)

### Small Messages

Data:
```
{
	"hello": "world",
	"name": "The Boss",
	"age": 29,
	"weight": 186.47
}
```
JSON: 60 bytes  
MessagePack: 48 bytes (22.2% smaller)

### Medium Messages

Data:
```
[
	{
		"hello": "world",
		"name": "The Boss",
		"age": 29,
		"weight": 186.47,
		"height": 72.3,
		"hobbies": ["hiking", "swimming", "reading"],
		"extra": {
			"location": "USA",
			"bio": "This \"string\" has \t characters \n that \" should \t be \\escaped."
		}
	},
	...repeat that X10
]
```
JSON: 2211 bytes  
MessagePack: 1841 bytes (18.3% smaller)

### Large messages

Data:
```
[
	{
		"hello": "world",
		"name": "The Boss",
		"age": 29,
		"weight": 186.47,
		"height": 72.3,
		"hobbies": ["hiking", "swimming", "reading", "soccer"],
		"extra": {
			"location": "USA",
			"bio": "This \n\t \"string\" has \n\tmore \t \"characters \n \" to \t \\escape.",
			"number": 105,
			"negative_num": -70
		}
	},
	...repeat that X30
]
```
JSON: 7951 bytes  
MessagePack: 6363 bytes (22.2% smaller)

#### Running the test

First, make sure you have Go installed. Then, after cloning the repo, from within the directory re-generate files for msgp in case the code generator has been updated:
```
go generate
```
Run the test:
```
go run *.go
```
