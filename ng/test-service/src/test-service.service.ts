import { Injectable } from '@angular/core';
import { NgClient } from '@1backend/ng-client';



export interface Test {
	name:	string;
	foods:	string[];
}


@Injectable()
export class Service {
  constructor(private ngClient: NgClient) {}

  GetHi(howManyTimes: number, allCool: boolean): Promise<void> {
    return this.ngClient.call<void>("dobika-test", "test-service", "GET", "/hi", { "howManyTimes": howManyTimes, "allCool": allCool });
  }

  GetImportedHi(): Promise<void> {
    return this.ngClient.call<void>("dobika-test", "test-service", "GET", "/imported-hi", {  });
  }

  GetSqlExample(): Promise<string> {
    return this.ngClient.call<string>("dobika-test", "test-service", "GET", "/sql-example", {  });
  }

}
