import { Component, OnInit } from '@angular/core';
import {CdkDragDrop, moveItemInArray, transferArrayItem} from '@angular/cdk/drag-drop';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-sprint-list',
  templateUrl: './sprint-list.component.html',
  styleUrls: ['./sprint-list.component.scss']
})
export class SprintListComponent implements OnInit {
  todo = ['Get to work', 'Pick up groceries', 'Go home', 'Fall asleep'];

  done = ['Get up', 'Brush teeth', 'Take a shower', 'Check e-mail', 'Walk dog'];
  sprintlist = [
    {'name': 'Sprint1', 'items':['Create delete button', 'Create Add button', 'Create Submit button']},
    {'name': 'Sprint2', 'items':['Create add API', 'Create Delete API', 'Write test cases']}
  ]
  
  constructor(private route: ActivatedRoute,private apiService: ApiInterfaceService) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));
    this.apiService.getSprintList({'data':projectIdFromRoute})
    .subscribe((resp:any) => {
      this.sprintlist = resp['response']
    })
  }

  drop(event: CdkDragDrop<string[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    } else {
      transferArrayItem(
        event.previousContainer.data,
        event.container.data,
        event.previousIndex,
        event.currentIndex,
      );
    }
  }

}
