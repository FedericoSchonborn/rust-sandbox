OPTION_LIBRARY = liboption.so
OPTION_OBJECTS = option.o

all: $(OPTION_LIBRARY)

$(OPTION_LIBRARY): $(OPTION_OBJECTS)
	$(CC) -shared -o $@ $<

%.o: %.c
	$(CC) -o $@ -c $<

clean:
	rm $(OPTION_LIBRARY) $(OPTION_OBJECTS)
