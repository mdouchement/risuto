import { Component, EventEmitter, Input, Output } from '@angular/core';

import { Item } from './item';
import { ItemService } from './item.service';

@Component({
  selector: 'risuto-new-item',
  template: `
    <div class="row" *ngIf="isActive">
      <div class="input-group">
        <span class="input-group-label">Name</span>
        <input [(ngModel)]="item.name" class="input-group-field" type="text" placeholder="Shingeki no Kyojin">
      </div>
      <div class="input-group">
        <span class="input-group-label">Description</span>
        <input [(ngModel)]="item.temporaryDescription" class="input-group-field" type="text" placeholder="A great anime!">
      </div>
      <div class="button-group tiny">
        <a class="button" (click)="goBack()">Back</a>
        <a class="button" (click)="save()">Save</a>
      </div>
    </div>
  `,
  providers: [ItemService]
})

export class NewItemComponent {
  item: Item;
  isActive: boolean;
  @Output() close = new EventEmitter();
  error: any;

  constructor(private itemService: ItemService) {}

  activate() {
    this.isActive = true;
    this.item = new Item();
    this.item.temporaryDescription = "";
  }

  save() {
    if (this.item.temporaryDescription != "") {
      this.item.descriptions.push(this.item.temporaryDescription);
    }

    this.itemService
        .save(this.item)
        .then(item => {
          this.item = item;
          this.goBack(item);
        })
        .catch(error => this.error = error);
  }

  goBack(item: item = null) {
    this.isActive = false;
    this.close.emit(item);
  }
}
