#include "isogram.h"
#include <stddef.h>
#include <ctype.h>
#include <stdio.h>

bool is_isogram(const char phrase[])
{
  int letters[27] = {0};

  if (phrase == NULL)
    return false;

  for (int i = 0; phrase[i] != '\0'; i++) {
    if (isalpha(phrase[i])) {
      int index = tolower(phrase[i]) - 'a';
      if (letters[index] == 1)
        return false;
      else
        letters[index] = 1;
    }
  }

  return true;
}
