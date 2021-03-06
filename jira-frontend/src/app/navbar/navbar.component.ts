import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiInterfaceService } from 'src/app/api-interface.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss']
})
export class NavbarComponent implements OnInit {

  constructor(private apiService:ApiInterfaceService, private router: Router) { }

  ngOnInit(): void {
  }

  logout(): void{
    this.apiService.removeToken();
    this.apiService.logout();
    this.router.navigateByUrl('login')
    // this.apiService.logout({})
    // .subscribe((resp:any) => {
    //   console.log(resp)
    //   this.apiService.setToken('')
    //   this.router.navigateByUrl('login')      
    // })
    
  }

}
