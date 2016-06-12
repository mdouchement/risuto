import { Pipe  } from '@angular/core';

@Pipe({
  name: "searchitem"
})

export class SearchItem {
  transform(items: Item[], itemFilter: string): Item[] {
    if (items !== undefined) {
      return items.filter(item => item.name.toLowerCase().search(itemFilter.toLowerCase()) != -1);
    }
    return items;
  }
}
