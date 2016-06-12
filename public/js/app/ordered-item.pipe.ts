import { Pipe  } from '@angular/core';

@Pipe({
  name: "ordereditem"
})

export class OrderedItem {
  transform(items: Item[]): Item[] {
    if (items !== undefined) {
      return items.sort((i1, i2) => i2.score - i1.score);
    }
    return items;
  }
}
