import { Injectable } from '@angular/core';
import { Headers, Http } from '@angular/http';

import 'rxjs/add/operator/toPromise';

import { Item } from './item';

@Injectable()
export class ItemService {
  private itemsUrl = pathJoin([namespace, '/items']);

  constructor(private http: Http) { }

  getItems(): Promise<Item[]>{
    return this.http.get(this.itemsUrl)
               .toPromise()
               .then(response => response.json())
               .catch(this.handleError);
  }

  save(item: Item): Promise<Item>  {
    if (item.id) {
      return this.update(item);
    }
    return this.create(item);
  }

  delete(item: Item) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');

    let url = `${this.itemsUrl}/${item.id}`;

    return this.http
               .delete(url, headers)
               .toPromise()
               .catch(this.handleError);
  }

  private create(item: Item): Promise<Item> {
    let headers = new Headers({
      'Content-Type': 'application/json'});

      return this.http
                 .post(this.itemsUrl, JSON.stringify(item), {headers: headers})
                 .toPromise()
                 .then(response => response.json())
                 .catch(this.handleError);
    }

  private update(item: Item) {
    let headers = new Headers();
    headers.append('Content-Type', 'application/json');

    let url = `${this.itemsUrl}/${item.id}`;

    return this.http
               .patch(url, JSON.stringify(item), {headers: headers})
               .toPromise()
               .then(() => item)
               .catch(this.handleError);
  }

  private handleError(error: any) {
    console.error('An error occurred', error);
    return Promise.reject(error.message || error);
  }
}
