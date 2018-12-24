# 练习使用命令
#exec:
#	cd /Users/liangliang;pwd
#

#ugn = haha
#foo = $(ugn)
#
#all:
#	echo $(foo)

#A = s
#B = $(A)
#
#all:
#	echo $(A)
#	echo $(B)

# 使用函数
#comm := ,
#conn := 
#space := $(conn) $(conn)
#foo := a b c
#bar := $(subst $(space),$(comm),$(foo))
#
#all:
#	echo $(bar)

# shell 函数
con := $(shell cat panic.go)

all:
	echo $(con)
