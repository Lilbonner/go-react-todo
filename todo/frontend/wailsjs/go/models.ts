export namespace main {
	
	export class Task {
	    id: number;
	    title: string;
	    dueDate: string;
	    dueTime: string;
	    priority: string;
	    completed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.dueDate = source["dueDate"];
	        this.dueTime = source["dueTime"];
	        this.priority = source["priority"];
	        this.completed = source["completed"];
	    }
	}

}

