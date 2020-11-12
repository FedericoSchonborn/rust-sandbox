OPTION_OBJECTS = option.o

all: liboption.so

liboption.so: $(OPTION_OBJECTS)
	$(CC) -shared -o $@ $<

%.o: %.c
	$(CC) -o $@ -c $<

clean:
	rm -f *.o *.so
