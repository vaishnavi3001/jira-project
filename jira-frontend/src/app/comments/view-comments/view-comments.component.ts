import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-view-comments',
  templateUrl: './view-comments.component.html',
  styleUrls: ['./view-comments.component.scss']
})
export class ViewCommentsComponent implements OnInit {

  constructor(private route: ActivatedRoute, private apiService:ApiInterfaceService) { }

  ngOnInit(): void {
    this.updatecomments();
  }

  sendcomment(comment:string): void {
    const routeParams = this.route.snapshot.paramMap;
    const issueIdFromRoute = Number(routeParams.get('issueId'));
    let body = {"issue_id":issueIdFromRoute,"comment":comment}
    console.log(body)
    this.apiService.createComment(body)
    .subscribe(res => {
          console.log(res);
    })
    this.updatecomments();
  }

  updatecomments(): void{
    const routeParams = this.route.snapshot.paramMap;
    const issueIdFromRoute = Number(routeParams.get('issueId'));
    let body = {"issue_id":issueIdFromRoute}
    console.log(body)
    this.apiService.getComments(body)
    .subscribe(res => {
          console.log(res['resp']['comments']);
    })
  }


}
