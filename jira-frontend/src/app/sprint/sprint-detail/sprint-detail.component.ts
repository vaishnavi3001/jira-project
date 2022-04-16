import { Component, OnInit } from '@angular/core';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { ActivatedRoute } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-sprint-detail',
  templateUrl: './sprint-detail.component.html',
  styleUrls: ['./sprint-detail.component.scss']
})
export class SprintDetailComponent implements OnInit {
  issue_list = []
  move_here :any[] = []
  constructor(private route: ActivatedRoute, private apiService:ApiInterfaceService) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const sprintIdFromRoute = Number(routeParams.get('sprintId'));
    console.log(sprintIdFromRoute)
       this.apiService.getIssueList({'user_id':1, 'sprint_id': sprintIdFromRoute})
       .subscribe((resp:any) =>{
       this.issue_list = (resp['resp']['issues'])
      })
       
  }

  Upcoming = [];
  InProgress = [];
  Review = [];
  Done = [];

  onDrop(event: CdkDragDrop<string[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    } else {
      transferArrayItem(event.previousContainer.data,
        event.container.data,
        event.previousIndex,
        event.currentIndex);
    }
  }  

}
