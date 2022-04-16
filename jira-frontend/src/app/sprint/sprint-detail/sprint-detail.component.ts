import { Component, OnInit } from '@angular/core';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-sprint-detail',
  templateUrl: './sprint-detail.component.html',
  styleUrls: ['./sprint-detail.component.scss']
})
export class SprintDetailComponent implements OnInit {
  issue_list :any[] = []
  move_here :any[] = []
  constructor(private route: ActivatedRoute, private apiService:ApiInterfaceService, private router: Router) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    const sprintIdFromRoute = Number(routeParams.get('sprintId'));
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

  goToIssueDetails(issueId: number): void{
    console.log(issueId)
    this.router.navigate(['home/issue/'+issueId+'/details']);
  }

  goToCreateIssue(): void{
    this.router.navigate(['home/'+1+'/issues/'+'/create']);
    
  }

}
