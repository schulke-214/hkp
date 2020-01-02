import { TaskClient } from './tasks';
import { TaskEntryClient } from './entries';

export class APIClient {
	constructor($axios) {
		this.tasks = new TaskClient($axios);
		this.entries = new TaskEntryClient($axios);
	}
}
