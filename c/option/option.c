#include <stdbool.h>
#include <stdlib.h>

#include "option.h"

struct Option
{
    bool ok;
    const void *value;
};

Option option_some(const void *value)
{
    return (Option){
        .ok = true,
        .value = value,
    };
}

Option option_none()
{
    return (Option){
        .ok = false,
        .value = NULL,
    };
}

Option option_map(Option self, OptionMapFunc func)
{
    if (!self.ok)
    {
        return option_none();
    }

    return func(self);
}
