import { Component, EventEmitter, Input, Output } from '@angular/core';

import { Item } from './item';
import { ItemService } from './item.service';

@Component({
  selector: 'risuto-item',
  styles: [`
    .focused-item {
      background-color: #fafbfb;
      border: 2px solid #cacaca;
      box-shadow: 0 2px #cacaca;
    }
  `],
  template: `
    <div class="callout" [class.focused-item]="item.isActive">
      <div class="row">
        <div class="columns">{{item.name}}</div>
        <div class="columns">
          <div class="float-right button-group tiny">
            <div class="secondary hollow label">{{item.score}}</div>
            <div class="shrink secondary hollow button tiny" (click)="dec()">-</div>
            <div class="shrink primary hollow button tiny" (click)="inc()">+</div>
          </div>
        </div>
      </div>
      <div [class.hide]="!item.isActive">
        <div class="row">
          <div class="columns">
            <ul *ngFor="let description of item.descriptions">
              <li>{{description}}</li>
            </ul>
          </div>
        </div>
        <div class="row">
          <div class="input-group">
            <input [(ngModel)]="item.temporaryDescription" class="input-group-field" type="text">
            <div class="input-group-button">
              <button class="button" (click)="addDescription()">+</button>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="columns"> </div>
          <div Class="float-right shrink alert hollow button tiny" (click)="delete(item, $event)">{{deleteMessages[deleteMessageIndex]}}</div>
        </div>
      </div>
    </div>
  `,
  providers: [ItemService],
  host: {
    class: 'column'
  }
})

export class ItemComponent {
  @Input() item: Item;
  deleteMessages = ["Delete", "Sure?"];
  deleteMessageIndex: int;
  @Output() deleted = new EventEmitter();
  error: any;

  constructor(private itemService: ItemService) {
    this.deleteMessageIndex = 0;
  }

  inc() {
    this.item.score++;
    this.itemService.save(this.item);
  }

  dec() {
    if (this.item.score > 0) {
      this.item.score--;
      this.itemService.save(this.item);
    }
  }

  addDescription() {
    this.item.descriptions.push(this.item.temporaryDescription);
    this.item.temporaryDescription = "";

    this.itemService
        .save(this.item)
        .then(item => {
          this.item = item;
          this.goBack(item);
        })
        .catch(error => this.error = error);
  }

  delete(item: Item, event: any) {
    if (this.deleteMessageIndex < 1) {
      this.deleteMessageIndex++;
      return;
    }

    event.stopPropagation();
    this.itemService
        .delete(item)
        .then(res => {
          this.deleted.emit(item);
        })
        .catch(error => this.error = error); // TODO: Display error message
  }
}
