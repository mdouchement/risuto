import { Component, OnInit } from '@angular/core';

import { Item } from './item';
import { ItemComponent } from './item.component';
import { ItemService } from './item.service';
import { OrderedItem } from './ordered-item.pipe';
import { SearchItem } from './search-item.pipe';

@Component({
  selector: 'risuto-items',
  template: `
    <!-- http://foundation.zurb.com/sites/docs/accordion.html -->
    <div class="row" *ngFor="let item of items | ordereditem | searchitem: itemFilter" (mouseenter)="item.isActive=true" (mouseleave)="item.isActive=false">
      <risuto-item [item]="item" (deleted)="deleted($event)"></risuto-item>
    </div>
  `,
  pipes: [OrderedItem, SearchItem],
  directives: [ItemComponent],
  providers: [ItemService]
})

export class ItemsComponent {
  items: Item[];
  itemFilter: string;
  error: any;

  constructor(private itemService: ItemService) {
    this.itemFilter = "";
  }

  getItems() {
    this.itemService
        .getItems()
        .then(items => this.items = items)
        .catch(error => this.error = error);
  }

  deleted(item: Item) {
    this.items = this.items.filter(i => i !== item);
  }

  ngOnInit() {
    this.getItems();
  }
}
