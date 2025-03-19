# frozen_string_literal: true

RSpec.describe GoGemYieldWithinGoroutine do
  describe ".sum" do
    subject { GoGemYieldWithinGoroutine.sum(a, b) }

    let(:a) { 1 }
    let(:b) { 2 }

    it { should eq 3 }
  end
end
