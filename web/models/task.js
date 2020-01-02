export class Task {
	constructor(task) {
		this.id = task.id;
		this.created_at = task.created_at;
		this.updated_at = task.updated_at;
		this.deleted_at = task.deleted_at;
		this.name = task.name;
	}
}
