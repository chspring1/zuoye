// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
	struct Voter{
		uint weight; // 投票权重
		bool voted; // 是否已投票
		address delegate; // 委托人
		uint vote; // 投票提案索引
	}
	struct Proposal {
		uint voteCount;
	}
	address public chairperson; //定义主席的地址
	Proposal[] public proposals; //定义提案数组
	mapping(address => Voter) public voters; //定义选民映射
	// 构造函数，初始化主席和提案
	constructor(uint proposalCount) {
		chairperson = msg.sender; //设置合约创建者为主席
		voters[chairperson].weight = 1; //主席的权重为1
		for (uint i = 0; i < proposalCount; i++) {
			proposals.push(Proposal(0));
		}
	}
	// 赋予投票权函数
	function giveRightToVote(address voter) public {
		require(msg.sender == chairperson, unicode"只有主席可以赋予投票权");
		require(!voters[voter].voted, unicode"选民已经投过票");
		require(voters[voter].weight == 0, unicode"选民已经有投票权");
		voters[voter].weight = 1; //赋予选民投票权
	}
//委托投票
function delegateVote(address to) public {
  Voter storage sender = voters[msg.sender];// 获取发送者的选民信息
  require(!sender.voted, unicode"你已经投过票了");
  require(to != msg.sender, unicode"不能委托给自己");
   while (voters[to].delegate != address(0) && voters[to].delegate != msg.sender) {
	to = voters[to].delegate; //找到最终的委托人
	}
	  require(to != msg.sender, unicode"不能委托给自己");
	  sender.voted = true; //标记选民已投票
	  sender.delegate = to; //设置委托人
	Voter storage delegateTo = voters[to];
	  if (delegateTo.voted) {
		proposals[delegateTo.vote].voteCount += sender.weight; //将选民的权重加到委托人的提案上
	}else {
		delegateTo.weight += sender.weight; //将选民的权重加到委托人身上
	}
}
// 投票
function vote(uint proposal) public {
	Voter storage sender = voters[msg.sender];// 获取发送者的选民信息
	require(sender.weight != 0, unicode"没有投票权");// 确保选民有投票权
	require(!sender.voted, unicode"已经投过票");// 确保选民没有投过票
	sender.voted = true; // 标记选民已投票
	sender.vote = proposal;// 记录选民的投票提案
	proposals[proposal].voteCount += sender.weight; // 将选民的权重加到提案上
}

// 查询获胜提案
function winningProposal() public view returns (uint winningProposal_) {
	uint winningVoteCount = 0; // 记录获胜提案的得票数
	for (uint i = 0; i < proposals.length; i++) {
		if (proposals[i].voteCount > winningVoteCount) {
			winningVoteCount = proposals[i].voteCount;
			winningProposal_ = i; //记录获胜提案的索引
		}
	}
	return winningProposal_;
}
}