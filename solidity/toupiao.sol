// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    struct Proposal {
        string name;      // 提案名称
        uint voteCount;   // 得票数
    }

    address public owner;
    mapping(address => bool) public hasVoted;
    Proposal[] public proposals;

    constructor(string[] memory proposalNames) {
        owner = msg.sender;
        for (uint i = 0; i < proposalNames.length; i++) {
            proposals.push(Proposal({name: proposalNames[i], voteCount: 0}));
        }
    }

    // 投票
    function vote(uint proposalIndex) public {
        require(!hasVoted[msg.sender], "You have already voted.");
        require(proposalIndex < proposals.length, "Invalid proposal index.");
        hasVoted[msg.sender] = true;
        proposals[proposalIndex].voteCount += 1;
    }

    // 查询所有提案
    function getProposals() public view returns (Proposal[] memory) {
        return proposals;
    }

    // 查询获胜提案
    function winningProposal() public view returns (string memory winnerName, uint winnerVotes) {
        uint maxVotes = 0;
        uint winnerIndex = 0;
        for (uint i = 0; i < proposals.length; i++) {
            if (proposals[i].voteCount > maxVotes) {
                maxVotes = proposals[i].voteCount;
                winnerIndex = i;
            }
        }
        winnerName = proposals[winnerIndex].name;
        winnerVotes = proposals[winnerIndex].voteCount;
    }
}
