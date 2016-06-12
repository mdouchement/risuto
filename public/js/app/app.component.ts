import { Component, ViewChild } from '@angular/core';
import { ItemsComponent } from './items.component';
import { NewItemComponent } from './new-item.component';

@Component({
  selector: 'risuto-app',
  styles: [`
    header {
      background-color: #d7d8d9;
      padding: .5rem;
      margin-bottom: 2rem;
    }
    header .row { position: relative; }
    header button {
      border-radius: 50%;
      height: 40px;
      position: absolute;
      right: 0;
      top: 30px;
      background-color: #e1683d;
      font-weight: bold;
    }
    header input {
      margin-bottom: 0;
      width: 90%;
    }
  `],
  template: `
    <header>
      <div class="row align-middle">
        <div class="column">
          <h4>Risuto</h4>
        </div>
        <div class="column">
          <button type="button" class="button" (click)="newItem.activate()">+</button>
          <input [(ngModel)]="items.itemFilter" type="search" placeholder="Search">
        </div>
      </div>
    </header>
    <risuto-new-item #newItem (close)="close($event)"></risuto-new-item>
    <risuto-items #items>Loading...</risuto-items>
  `,
  directives: [ItemsComponent, NewItemComponent]
})

export class AppComponent {
  @ViewChild(ItemsComponent) items:ItemsComponent; // Child accessor

  close(item: Item) {
    if (item) { this.items.getItems(); }
  }
}
