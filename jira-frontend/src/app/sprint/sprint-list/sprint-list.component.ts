import { Component, OnInit } from '@angular/core';
import {CdkDragDrop, moveItemInArray, transferArrayItem} from '@angular/cdk/drag-drop';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { ActivatedRoute } from '@angular/router';

export interface SprintDetails {
  id: number,
  name: string;
  project_id: number;
  created_at: string;
  start_date: string;
  end_date: string;
}

@Component({
  selector: 'app-sprint-list',
  templateUrl: './sprint-list.component.html',
  styleUrls: ['./sprint-list.component.scss']
})
export class SprintListComponent implements OnInit {


  todo = ['Get to work', 'Pick up groceries', 'Go home', 'Fall asleep'];

  done = ['Get up', 'Brush teeth', 'Take a shower', 'Check e-mail', 'Walk dog'];

  constructor(private route: ActivatedRoute,private apiService: ApiInterfaceService) { }

  sprintlist: SprintDetails[] = [
  ];

  issueList : any[] | any
  mainIssueList : any[] | any

  ngOnInit(): void { 

    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'));
    this.apiService
    .getSprintList({"project_id":projectIdFromRoute, "user_id":1 })
    .subscribe((resp:any) => {
      this.sprintlist = (resp['resp']['sprints'])
      this.sprintlist.forEach( (value: { name: any; id: any; }) => {
        this.apiService.getIssueList({"user_id":1, "sprint_id":value.id})
        .subscribe((resp:any) => {
          this.issueList = resp
          this.mainIssueList = this.issueList['resp']['issues']
        })
      })
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
