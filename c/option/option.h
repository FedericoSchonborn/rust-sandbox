#pragma once

typedef struct Option Option;

Option Some(const void *value);
Option None();

typedef Option (*OptionMapFunc)(Option);

Option Map(Option self, OptionMapFunc func);
