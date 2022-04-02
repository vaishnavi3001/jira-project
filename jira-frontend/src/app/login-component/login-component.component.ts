import { Component, OnInit } from '@angular/core';

import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-login-component',
  templateUrl: './login-component.component.html',
  styleUrls: ['./login-component.component.scss']
})
export class LoginComponentComponent implements OnInit {
  
  newLoginForm = new FormGroup({
    username: new FormControl('', [Validators.required, Validators.minLength(3)]),
    password: new FormControl('', [Validators.required]),
  });

  constructor(private apiService: ApiInterfaceService, private route: ActivatedRoute, private router: Router ) { }

  ngOnInit(): void {
  }

  submit(){
    let body = {
      username:this.newLoginForm.get('username')?.value,
      password: this.newLoginForm.get('password')?.value,
    } 

    this.apiService.login(body)
    .subscribe(res => {
          console.log(res);
          this.router.navigate(['home/project/list']);
    })
    
  }

}
