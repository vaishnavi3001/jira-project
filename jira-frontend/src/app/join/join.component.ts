import { ThrowStmt } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-join',
  templateUrl: './join.component.html',
  styleUrls: ['./join.component.scss']
})
export class JoinComponent implements OnInit {

  constructor(private route: ActivatedRoute, private router:Router) {

  }

  ngOnInit(): void {
    let joinlink:string = this.route.snapshot.params['join-link']
    if (joinlink && joinlink.length!=0){
      
    }
    else{
      this.router.navigate(['home']);
    }
  }


  checkLink(): void{
    
  }


  




}
