# frozen_string_literal: true

RSpec.describe GoGemYieldWithinGoroutine do
  describe ".sum" do
    subject { GoGemYieldWithinGoroutine.sum(a, b) }

    let(:a) { 1 }
    let(:b) { 2 }

    it { should eq 3 }
  end

  describe ".with_block" do
    context "with block" do
      it "works" do
        actual =
          GoGemYieldWithinGoroutine.with_block(2) do |a|
            a * 3
          end

        expect(actual).to eq 6
      end
    end

    context "without block" do
      it { expect { GoGemYieldWithinGoroutine.with_block(2) }.to raise_error ArgumentError }
    end
  end
end
