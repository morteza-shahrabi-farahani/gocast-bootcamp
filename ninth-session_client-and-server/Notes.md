\* If the fields of the request and response are different from the struct type, it is better to seperate request and response and entity types.

\* The current design and current types would meet your current needs; However, in the future, it would be probable that some part of the design will change. So it is better for now to consider these probable future changes and change them now in your code. Because if you don't do it now and postpone it to the future, the cost of change will grow exponentially and become too harder. You may notice that this point is contradicts the YAGNI principle. YAGNI developing principle means You aren't gonna need it. This is relate to some scenarios which you make the code and say that in the future it will change but you don't need it now.

The important point that you should consider in all the time when you are developing a software is that as a software engineer, one of the most priority of your work after its functionality and working correctly, is that you should minimum the required cost for change in the future. So in naming the variables, in making the function signitures, in writing or not writing interfaces and all the time, you should say with your self that as far as cost of change is concerned, I must do this. I should take this decision. So in one scenario if you find that some of the principles are in contrast with this important fact, throw away those principles. Principles are not correct in all times.

\* In service layer, we handle the logic of our program. So we should not now about the storage of the program and the repository layer and there functionalities and behaviors must be handled in their layer separately.

One way for handling this is to have a repository interface in our service part and in this interface we write the signiture of the functions that we need repository to handle them. The implemantations of these functionalities are in the repository package.

\* If our request part just contains one field, but it is better to make a new type for your request for maintaining integrity of your codes and your project.
