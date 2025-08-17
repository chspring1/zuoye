// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BeggingContract {
    struct donorDetail{
        address donor; //捐赠者
        uint amount;  //捐赠数
        uint timestamp; //时间戳
    }
    address[] public donors;
    mapping (address => donorDetail) public sessions;
    uint public totalDonations;
    address public immutable constantowner;

    //初始化执行

    constructor() {
        constantowner = msg.sender; //记录合约创建者
    }
    //捐赠函数
    function donate ()public payable {
        require(msg.value >0 ,"The donation amount must be greater than 0");
        sessions[msg.sender] = donorDetail(msg.sender,msg.value,block.timestamp);
        totalDonations +=msg.value;
        donors.push(msg.sender);
    }
    //取款函数
    function withDraw(uint amount)public {
        require(msg.sender == constantowner,"You do not have the authority to operate");//只有合约拥有者可以取款
        require(totalDonations >0,"Insufficient Balance");//校验余额
        totalDonations -=amount ;
        payable(msg.sender).transfer(amount); //将当前账户的钱转给合约调用者
    }
    //查询捐赠者信息
    function getDonor (address donor_) public view returns (uint amount,uint timestamp){
        donorDetail memory session = sessions[donor_];
        return (session.amount,session.timestamp);
    }

    //查询自己的捐赠信息
    
    function getOwnInformation() public view returns (uint amount,uint timestamp){
       bool memory OK = false;
    for (uint i= 0; i< donors.length;i++){
            if(donors[i]==msg.sender){
                OK = true ;
                break ;
            }
       }
        require(OK,"No record");
        donorDetail memory session = sessions[msg.sender];
         return (session.amount,session.timestamp);
    }
  
   //查询余额

   function getTotalDonations()public view returns (uint TotalDonations_){
            return totalDonations;
        }   
   
 }
    
    
