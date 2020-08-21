import nltk
from sys import stdin

spanR = "<span class=\"regular\">"
spanI = "<span class=\"irregular\">"

input_data = input() 

morph = nltk.word_tokenize(input_data)
pos = nltk.pos_tag(morph)

for i in pos:
	if i[1] == "VBD":
		if i[0][len(i[0]) - 2:] == 'ed':
			print(spanR + i[0] + "</span>", end=" ")
		else:
			print(spanI + i[0] + "</span>", end=" ")
	else:
		print(i[0], end=" ")

