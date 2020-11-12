#pragma once

typedef struct Option Option;

Option option_some(const void *value);
Option option_none();

typedef Option (*OptionMapFunc)(Option);

Option option_map(Option self, OptionMapFunc func);
