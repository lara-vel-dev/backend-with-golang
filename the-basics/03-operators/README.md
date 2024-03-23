<h1>Operators</h1>
<p align="justify"> 
Operators, like in many other programming languages, perform various operations and act like tools that manipulate dara and control the flow of a program. These are classified into different categories based on their functionality.
</p>

<div>
	<h3>Arithmetic Operators</h3>
	<p align="justify"> 
	Perform basic mathematical calculations, such as addition, substraction, multiplication, division and finding the remainder. These are used for numerical computations, manipulating numeric variables and performing calculations within expressions.
	</p>
	<table border='1' align="center"> 
		<thead>
			<tr align='center'>
				<td> <b>Operator</b> </td>
				<td> <b>Symbol</b> </td>
			</tr>
		</thead>
		<tbody>
			<tr align='center'>
				<td>Addition</td>
				<td>+</td>
			</tr>
			<tr align='center'>
				<td>Subtraction</td>
				<td>-</td>
			</tr>
			<tr align='center'>
				<td>Multiplication</td>
				<td>*</td>
			</tr>
			<tr align='center'>
				<td>Division</td>
				<td>%</td>
			</tr>
			<tr align='center'>
				<td>Remainder/Module</td>
				<td>+</td>
			</tr>
		</tbody>
	</table>
</div>

<div>
	<h3>Relational Operators</h3>
	<p align="justify"> 
	Compare values and return boolean results (true or false) based on the comparison. The comparison values can be numbers, characters, strings.
	</p>
	<ul>
		<li><code>==</code> : Checks if both operands have the same value.</li>
		<li><code>!=</code> : Checks if both operands have different values.</li>
		<li><code>>=</code> : Checks if the left operand is greater than or equal to the right operand.</li>
		<li><code><=</code> : Checks if the left operand is less than or equal to the right operand.</li>
		<li><code>></code> : Checks if the left operand is greater than the right operand.</li>
		<li><code><</code> : Checks if the left operand is less than the right operand.</li>
	</ul>
	<table border='1' align="center"> 
		<thead>
			<tr align='center'>
				<td> <b>Operator</b> </td>
				<td> <b>Symbol</b> </td>
			</tr>
		</thead>
		<tbody>
			<tr align='center'>
				<td>Equal to</td>
				<td>==</td>
			</tr>
			<tr align='center'>
				<td>Not equal to</td>
				<td>!=</td>
			</tr>
			<tr align='center'>
				<td>Greater than or equal to</td>
				<td>>=</td>
			</tr>
			<tr align='center'>
				<td>Less than or equal to</td>
				<td><=</td>
			</tr>
			<tr align='center'>
				<td>Greater than</td>
				<td>></td>
			</tr>
			<tr align='center'>
				<td>Less than</td>
				<td><</td>
			</tr>
		</tbody>
	</table>
</div>

<div>
	<h3>Logical Operators</h3>
	<p align="justify"> 
	Combine conditional expressions and control the flow of the program. They operate on boolean values (true or false) and return a boolean result based on the operands and the specific logical operator used.
	</p>
		<ul>
		<li><code>&&</code> : Represents a conjunction, both operands must be true for the result to be true.</li>
		<li><code>||</code> : Represents a disjunction, at least one operand must be true for the result ti be true.</li>
		<li><code>!</code> : Inverts the logical state of the operand, true becomes false and false becomes true.</li>
	</ul>
	<table border='1' align='center'> 
		<thead>
			<tr align='center'>
				<td> <b>Operator</b> </td>
				<td> <b>Symbol</b> </td>
			</tr>
		</thead>
		<tbody>
			<tr align='center'>
				<td>AND</td>
				<td>&&</td>
			</tr>
			<tr align='center'>
				<td>OR</td>
				<td>||</td>
			</tr>
			<tr align='center'>
				<td>NOT</td>
				<td>!</td>
			</tr>
		</tbody>
	</table>
	<table border='1' align='center'> 
		<thead>
			<tr align='center'>
				<td colspan="2"> <b>AND Truth Table</b> </td>
			</tr>
			<tr align='center'>
				<td> <b>Expression</b> </td>
				<td><b>Value</b></td>
			</tr>
		</thead>
		<tbody>
			<tr align='center' align='center'>
				<td><code>true && true</code></td>
				<td><code>true</code></td>
			</tr>
			<tr align='center'>
				<td><code>true && false</code></td>
				<td><code>false</code></td>
			</tr>
			<tr align='center'>
				<td><code>false && true</code></td>
				<td><code>false</code></td>
			</tr>
			<tr align='center'>
				<td><code>false && false</code></td>
				<td><code>false</code></td>
			</tr>
		</tbody>
	</table>
	<table border='1' align='center'> 
		<thead>
			<tr align='center'>
				<td colspan="2"> <b>OR Truth Table</b> </td>
			</tr>
			<tr align='center'>
				<td> <b>Expression</b> </td>
				<td><b>Value</b></td>
			</tr>
		</thead>
		<tbody>
			<tr align='center'>
				<td><code>true || true</code></td>
				<td><code>true</code></td>
			</tr>
			<tr align='center'>
				<td><code>true || false</code></td>
				<td><code>true</code></td>
			</tr>
			<tr align='center'>
				<td><code>false || true</code></td>
				<td><code>true</code></td>
			</tr>
			<tr align='center'>
				<td><code>false || false</code></td>
				<td><code>false</code></td>
			</tr>
		</tbody>
	</table>
	<table border='1' align='center'> 
		<thead>
			<tr align='center'>
				<td colspan="2"> <b>NOT Truth Table</b> </td>
			</tr>
			<tr align='center'>
				<td> <b>Expression</b> </td>
				<td><b>Value</b></td>
			</tr>
		</thead>
		<tbody>
			<tr align='center'>
				<td><code>!true</code></td>
				<td><code>false</code></td>
			</tr>
			<tr align='center'>
				<td><code>!false</code></td>
				<td><code>true</code></td>
			</tr>
		</tbody>
	</table>
</div>

<div>
	<h3>Bitwise Operators</h3>
	<p align="justify"> 
	Manipulate data at bit level. They operate on individual bits within integers, providing control over the binary representation of a number.
	</p>
		<ul>
		<li><code>>></code> : Shifts the bits in the left operand to the right by the number of the position specified by the right operand.</li>
		<li><code><<</code> : Shifts the bits in the left operand to the left by the number of the position specified by the right operand.</li>
		<li><code>&</code> : Performs a bitwise AND operation on each corresponding bit of the operands.</li>
		<li><code>|</code> : Performs a bitwise OR operation on each corresponding bit of the operands.</li>
		<li><code>^</code> : Performs a bitwise XOR operation on each corresponding bit of the operands.</li>
	</ul>
	<table border='1' align='center'> 
		<thead>
			<tr align='center'>
				<td> <b>Operator</b> </td>
				<td> <b>Symbol</b> </td>
			</tr>
		</thead>
		<tbody>
			<tr align='center'>
				<td>Right shift</td>
				<td>>></td>
			</tr>
			<tr align='center'>
				<td>Left shift</td>
				<td><<</td>
			</tr>
			<tr align='center'>
				<td>AND</td>
				<td>&</td>
			</tr>
			<tr align='center'>
				<td>OR</td>
				<td>|</td>
			</tr>
			<tr align='center'>
				<td>XOR</td>
				<td>^</td>
			</tr>
		</tbody>
	</table>
</div>

<div>
	<h3>Miscellaneous Operators</h3>
	<p align="justify"> 
	Provide various functionalities for memory management, data type manipulation, and shorthand notations. 
	</p>
		<ul>
		<li><code>&</code> : Returns the memory address of a variable.</li>
		<li><code>*</code> : Retrieves the value stored at the memory address pointed to by a pointer.</li>
		<li><code>++</code> : Provide shorthand notations for incrementing a variable's value by 1.</li>
		<li><code>--</code> : Provide shorthand notations for decrementing a variable's value by 1.</li>
	</ul>
	<table border='1' align='center'> 
		<thead>
			<tr align='center'>
				<td> <b>Operator</b> </td>
				<td> <b>Symbol</b> </td>
			</tr>
		</thead>
		<tbody>
			<tr align='center'>
				<td>Address of</td>
				<td>&</td>
			</tr>
			<tr align='center'>
				<td>Dereference</td>
				<td>*</td>
			</tr>
			<tr align='center'>
				<td>Increment</td>
				<td>++</td>
			</tr>
			<tr align='center'>
				<td>Decrement</td>
				<td>--</td>
			</tr>
		</tbody>
	</table>
</div>

<div>
	<h3>Assignment Operators</h3>
	<p align="justify"> 
	Assign values to variables but also offer several compound assignment operators that combine an arithmetic operation with assignment.
	</p>
		<ul>
		<li><code>=</code> : It assigns the value of the expression on the right side to the variable on the left side.
</li>
		<li><code>:=</code> : Declares a new variable and assigns a value to the newly declared variable.</li>
		<li><code>+=</code> : Adds the value on the right to the variable on the left and assigns the result back to the variable.</li>
		<li><code>-=</code> : Subtracts the value on the right from the variable on the left and assigns the result back to the variable.</li>
		<li><code>*=</code> : Multiplies the variable on the left by the value on the right and assigns the result back to the variable.</li>
		<li><code>/=</code> : Divides the variable on the left by the value on the right and assigns the result back to the variable.</li>
		<li><code>%=</code> : Computes the remainder of dividing the variable on the left by the value on the right and assigns the result back to the variable.</li>
		<li><code>>>=</code> : Performs a right shift operation on the variable on the left by the number of bits specified on the right and assigns the result back to the variable.</li>
		<li><code><<=</code> : Performs a left shift operation on the variable on the left by the number of bits specified on the right and assigns the result back to the variable.</li>
		<li><code>&=</code> : Performs a bitwise AND operation between the variable on the left and the value on the right and assigns the result back to the variable.</li>
		<li><code>|=</code> : Performs a bitwise OR operation between the variable on the left and the value on the right and assigns the result back to the variable.</li>
		<li><code>^=</code> : Performs a bitwise XOR operation between the variable on the left and the value on the right and assigns the result back to the variable.</li>
	</ul>
	<table border='1' align='center'> 
	<thead>
		<tr align='center'>
			<td> <b>Operator</b> </td>
			<td> <b>Symbol</b> </td>
		</tr>
	</thead>
	<tbody>
		<tr align='center'>
			<td>Assign</td>
			<td>=</td>
		</tr>
		<tr align='center'>
			<td>Declare and assign</td>
			<td>:=</td>
		</tr>
		<tr align='center'>
			<td>Add and assign</td>
			<td>+=</td>
		</tr>
		<tr align='center'>
			<td>Subtract and assign</td>
			<td>-=</td>
		</tr>
		<tr align='center'>
			<td>Multiply and assign</td>
			<td>*=</td>
		</tr>
		<tr align='center'>
			<td>Divide and assign</td>
			<td>/=</td>
		</tr>
		<tr align='center'>
			<td>Module assignment</td>
			<td>%=</td>
		</tr>
		<tr align='center'>
			<td>Right shift assignment</td>
			<td>>>=</td>
		</tr>
		<tr align='center'>
			<td>Left shift assignment</td>
			<td><<=</td>
		</tr>
		<tr align='center'>
			<td>Bitwise AND assignment</td>
			<td>&=</td>
		</tr>
		<tr align='center'>
			<td>Bitwise OR assignment</td>
			<td>|=</td>
		</tr>
		<tr align='center'>
			<td>Bitwise XOR assignment</td>
			<td>^=</td>
		</tr>
	</tbody>
</table>

</div>

<div>
<a href="https://github.com/lara-vel-dev/backend-with-golang/tree/main/the-basics/02-variables-and-data-types" >
	<strong>⬅️ Previous topic</strong>
</a>
&emsp;
<a href="https://github.com/lara-vel-dev/backend-with-golang/blob/main/the-basics/03-operators" >
	<strong>➡️ Next topic</strong>
</a>
</div>
