#include <stdbool.h>
#include <stdlib.h>

#include "option.h"

struct Option
{
    bool ok;
    const void *value;
};

Option Some(const void *value)
{
    return (Option){
        .ok = true,
        .value = value,
    };
}

Option None()
{
    return (Option){
        .ok = false,
        .value = NULL,
    };
}

Option Map(Option self, OptionMapFunc func)
{
    if (!self.ok)
    {
        return None();
    }

    return func(self);
}
