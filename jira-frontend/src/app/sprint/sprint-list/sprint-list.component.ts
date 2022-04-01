import { Component, OnInit } from '@angular/core';
import {CdkDragDrop, moveItemInArray, transferArrayItem} from '@angular/cdk/drag-drop';
import { ApiInterfaceService } from 'src/app/api-interface.service';
import { ActivatedRoute } from '@angular/router';
import { Router } from '@angular/router';

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

  constructor(private route: ActivatedRoute,private apiService: ApiInterfaceService, private router: Router) { }

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
    })
  }

  openCreateSprint(): void{
    const routeParams = this.route.snapshot.paramMap;
    const projectIdFromRoute = Number(routeParams.get('projectId'))
    let _route = 'home/'+projectIdFromRoute+'/sprint/create'
    this.router.navigateByUrl(_route);
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
